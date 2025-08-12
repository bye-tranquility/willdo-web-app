import { useState } from 'react'
import EventItem from './EventItem'

const EventList = ({ events, loading, onToggleComplete, onUpdate, onDelete }) => {
  const [filter, setFilter] = useState('all') // 'all', 'pending', 'completed'

  if (loading) {
    return (
      <div className="card">
        <div className="loading">Loading your tasks...</div>
      </div>
    )
  }

  const filteredEvents = events.filter(event => {
    if (filter === 'pending') return !event.completed
    if (filter === 'completed') return event.completed
    return true
  })

  const pendingCount = events.filter(event => !event.completed).length
  const completedCount = events.filter(event => event.completed).length

  return (
    <div className="card event-list">
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '1.5rem' }}>
        <h2>Your Tasks</h2>
        <div style={{ display: 'flex', gap: '0.5rem' }}>
          <button
            className={`btn btn-small ${filter === 'all' ? 'btn-primary' : 'btn-secondary'}`}
            onClick={() => setFilter('all')}
          >
            All ({events.length})
          </button>
          <button
            className={`btn btn-small ${filter === 'pending' ? 'btn-primary' : 'btn-secondary'}`}
            onClick={() => setFilter('pending')}
          >
            Pending ({pendingCount})
          </button>
          <button
            className={`btn btn-small ${filter === 'completed' ? 'btn-primary' : 'btn-secondary'}`}
            onClick={() => setFilter('completed')}
          >
            Completed ({completedCount})
          </button>
        </div>
      </div>

      {filteredEvents.length === 0 ? (
        <div className="empty-state">
          <h3>
            {filter === 'all'
              ? "No tasks yet"
              : filter === 'pending'
                ? "No pending tasks"
                : "No completed tasks"}
          </h3>
          <p>
            {filter === 'all'
              ? "Add your first task above to get started!"
              : filter === 'pending'
                ? "Great job! All tasks are completed."
                : "Complete some tasks to see them here."}
          </p>
        </div>
      ) : (
        <div className="events-container">
          {filteredEvents.map(event => (
            <EventItem
              key={event.id}
              event={event}
              onToggleComplete={onToggleComplete}
              onUpdate={onUpdate}
              onDelete={onDelete}
            />
          ))}
        </div>
      )}
    </div>
  )
}

export default EventList
