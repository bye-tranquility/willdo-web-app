import { useState } from 'react'

const EventForm = ({ onSubmit }) => {
  const [formData, setFormData] = useState({
    description: '',
    due: '',
    completed: false
  })
  const [isSubmitting, setIsSubmitting] = useState(false)

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target
    setFormData(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }))
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    if (!formData.description.trim()) {
      alert('Please enter a description for your task')
      return
    }

    setIsSubmitting(true)
    try {
      await onSubmit(formData)
      setFormData({
        description: '',
        due: '',
        completed: false
      })
    } catch (error) {
      console.error('Error submitting form:', error)
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <div className="card event-form">
      <h2>Add New Task</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="description">Task Description *</label>
          <textarea
            id="description"
            name="description"
            value={formData.description}
            onChange={handleChange}
            placeholder="What do you need to do?"
            required
          />
        </div>
        
        <div className="form-group">
          <label htmlFor="due">Due Date & Time</label>
          <input
            type="datetime-local"
            id="due"
            name="due"
            value={formData.due}
            onChange={handleChange}
          />
        </div>

        <button 
          type="submit" 
          className="btn btn-primary"
          disabled={isSubmitting}
        >
          {isSubmitting ? 'Adding Task...' : 'Add Task'}
        </button>
      </form>
    </div>
  )
}

export default EventForm
