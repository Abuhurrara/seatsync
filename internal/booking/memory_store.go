package booking

type MemoryStore struct {
	bookings map[string]Booking
}

// constructor
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: make(map[string]Booking),
	}
}

func (s *MemoryStore) Book(b Booking) (Booking, error) {
	if _, exists := s.bookings[b.SeatID]; exists {
		return b, ErrSeatAlreadyExists
	}
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {

}
