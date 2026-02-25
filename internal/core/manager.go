package core

import "github.com/scvrylullaby/bowling-centre-backend/internal/models"

type Manager struct {
	Incoming  chan *models.Client
	LaneFreed chan int
	LeftQueue chan int
	StateChan chan models.DashboardState

	poolSize    int
	queue       []*models.Client
	freeLanes   []int
	activeLanes map[int]int
	lanesChans  map[int]chan *models.Client
	stats       models.Stats
}

func NewManager(poolSize int, stateChan chan models.DashboardState) *Manager {
	m := &Manager{
		Incoming:    make(chan *models.Client),
		LaneFreed:   make(chan int),
		LeftQueue:   make(chan int),
		StateChan:   stateChan,
		poolSize:    poolSize,
		queue:       make([]*models.Client, 0),
		freeLanes:   make([]int, 0, poolSize),
		activeLanes: make(map[int]int),
		lanesChans:  make(map[int]chan *models.Client),
	}

	for i := 1; i <= poolSize; i++ {
		m.freeLanes = append(m.freeLanes, i)
		m.lanesChans[i] = make(chan *models.Client)
		go RunLane(i, m.lanesChans[i], m.LaneFreed)
	}

	return m
}

func (m *Manager) Run() {
	for {
		select {
		case clClient := <-m.Incoming:
			if len(m.freeLanes) > 0 {
				m.assignLane(clClient)
			} else {
				m.queue = append(m.queue, clClient)
				m.stats.Waiting++
			}

		case laneID := <-m.LaneFreed:
			m.stats.Playing--
			m.stats.Finished++
			delete(m.activeLanes, laneID)
			m.freeLanes = append(m.freeLanes, laneID)

			m.processQueue()

		case clClientID := <-m.LeftQueue:
			m.removeFromQueue(clClientID)
		}

		m.broadcastState()
	}
}

func (m *Manager) assignLane(c *models.Client) {
	laneID := m.freeLanes[0]
	m.freeLanes = m.freeLanes[1:]

	m.activeLanes[laneID] = c.ID
	m.stats.Playing++

	close(c.Start)
	m.lanesChans[laneID] <- c
}

func (m *Manager) processQueue() {
	if len(m.queue) > 0 && len(m.freeLanes) > 0 {
		nextClient := m.queue[0]
		m.queue = m.queue[1:]
		m.stats.Waiting--
		m.assignLane(nextClient)
	}
}

func (m *Manager) removeFromQueue(clClientID int) {
	for i, c := range m.queue {
		if c.ID == clClientID {
			m.queue = append(m.queue[:i], m.queue[i+1:]...)
			m.stats.Waiting--
			m.stats.Left++
			break
		}
	}
}

func (m *Manager) broadcastState() {
	activeCopy := make(map[int]int)
	for k, v := range m.activeLanes {
		activeCopy[k] = v
	}

	m.StateChan <- models.DashboardState{
		ActiveLanes: activeCopy,
		Stats:       m.stats,
	}
}
