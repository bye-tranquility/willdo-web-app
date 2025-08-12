import { useState } from 'react'

const EventItem = ({ event, onToggleComplete, onUpdate, onDelete }) => {
  const [isEditing, setIsEditing] = useState(false)
  const [editFormData, setEditFormData] = useState({
    description: event.description,
    due: event.due || ''
  })

  const handleEditChange = (e) => {
    const { name, value } = e.target
    setEditFormData(prev => ({
      ...prev,
      [name]: value
    }))
  }

  const handleEditSubmit = async (e) => {
    e.preventDefault()
    if (!editFormData.description.trim()) {
      alert('Description cannot be empty')
      return
    }

    try {
      await onUpdate(event.id, {
        description: editFormData.description,
        due: editFormData.due
      })
      setIsEditing(false)
    } catch (error) {
      console.error('Error updating event:', error)
    }
  }

  const handleEditCancel = () => {
    setEditFormData({
      description: event.description,
      due: event.due || ''
    })
    setIsEditing(false)
  }

  const formatDate = (dateString) => {
    if (!dateString) return null
    try {
      return new Date(dateString).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      })
    } catch {
      return dateString
    }
  }

  const isOverdue = () => {
    if (!event.due || event.completed) return false
    return new Date(event.due) < new Date()
  }

  return (
    <div className={`event-item ${event.completed ? 'completed' : ''}`}>
      {isEditing ? (
        <form onSubmit={handleEditSubmit}>
          <div className="form-group">
            <textarea
              name="description"
              value={editFormData.description}
              onChange={handleEditChange}
              placeholder="Task description"
              required
              style={{ marginBottom: '1rem' }}
            />
          </div>
          <div className="form-group">
            <input
              type="datetime-local"
              name="due"
              value={editFormData.due}
              onChange={handleEditChange}
              style={{ marginBottom: '1rem' }}
            />
          </div>
          <div style={{ display: 'flex', gap: '0.5rem' }}>
            <button type="submit" className="btn btn-small btn-success">
              Save
            </button>
            <button
              type="button"
              className="btn btn-small btn-secondary"
              onClick={handleEditCancel}
            >
              Cancel
            </button>
          </div>
        </form>
      ) : (
        <>
          <div className="event-header">
            <div className="event-description">
              {event.description}
            </div>
            <div className="event-actions">
              <button
                className={`btn btn-small ${event.completed ? 'btn-secondary' : 'btn-success'}`}
                onClick={() => onToggleComplete(event)}
                title={event.completed ? 'Mark as pending' : 'Mark as completed'}
              >
                {event.completed ? 'Undo ↺' : 'Done ✓'}
              </button>
              <button
                className="btn btn-small btn-secondary"
                onClick={() => setIsEditing(true)}
                title="Edit task"
              >
                Edit
              </button>
              <button
                className="btn btn-small btn-danger"
                onClick={() => {
                  if (window.confirm('Are you sure you want to delete this task?')) {
                    onDelete(event.id)
                  }
                }}
                title="Delete task"
              >
                Delete
              </button>
            </div>
          </div>

          <div className="event-meta">
            <div>
              {event.due && (
                <span
                  className="event-due-date"
                  style={{
                    color: isOverdue() ? 'var(--danger-color)' : 'var(--text-secondary)',
                    fontWeight: isOverdue() ? '600' : '500'
                  }}
                >
                  Due: {formatDate(event.due)}
                  {isOverdue() && ' (Overdue)'}
                </span>
              )}
            </div>
            <div>
              <span className={`event-status ${event.completed ? 'status-completed' : 'status-pending'}`}>
                {event.completed ? 'Completed' : 'Pending'}
              </span>
            </div>
          </div>
        </>
      )}
    </div>
  )
}

export default EventItem
