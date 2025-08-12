import { useState, useEffect } from 'react'
import './styles/globals.css'
import './styles/variables.css'
import './styles/buttons.css'
import './styles/forms.css'
import './styles/header.css'
import './styles/events.css'
import './styles/card.css'
import './styles/states.css'
import './styles/responsive.css'
import EventList from './components/EventList'
import EventForm from './components/EventForm'
import Header from './components/Header'
import { API_BASE_URL } from './constants/AppConstants'

function App() {
  const [events, setEvents] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  const fetchEvents = async () => {
    try {
      setLoading(true)
      setError(null)
      const response = await fetch(`${API_BASE_URL}/events`)
      if (!response.ok) {
        throw new Error('Failed to fetch events')
      }
      const data = await response.json()
      setEvents(data || [])
    } catch (err) {
      setError(err.message)
      console.error('Error fetching events:', err)
    } finally {
      setLoading(false)
    }
  }

  const createEvent = async (eventData) => {
    try {
      const response = await fetch(`${API_BASE_URL}/events`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(eventData),
      })
      if (!response.ok) {
        throw new Error('Failed to create event')
      }

      await fetchEvents()
    } catch (err) {
      setError(err.message)
      console.error('Error creating event:', err)
    }
  }

  const updateEvent = async (id, eventData) => {
    try {
      const response = await fetch(`${API_BASE_URL}/events/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(eventData),
      })
      if (!response.ok) {
        throw new Error('Failed to update event')
      }

      await fetchEvents()
    } catch (err) {
      setError(err.message)
      console.error('Error updating event:', err)
    }
  }

  const deleteEvent = async (id) => {
    try {
      const response = await fetch(`${API_BASE_URL}/events/${id}`, {
        method: 'DELETE',
      })
      if (!response.ok) {
        throw new Error('Failed to delete event')
      }

      await fetchEvents()
    } catch (err) {
      setError(err.message)
      console.error('Error deleting event:', err)
    }
  }

  const toggleEventCompletion = async (event) => {
    await updateEvent(event.id, {
      description: event.description,
      due: event.due || "",
      completed: !event.completed
    })
  }

  useEffect(() => {
    fetchEvents()
  }, [])

  return (
    <div className="app">
      <Header />
      <div className="container">
        <div className="main-content">
          <EventForm onSubmit={createEvent} />
          {error && <div className="error-message">Error: {error}</div>}
          <EventList 
            events={events}
            loading={loading}
            onToggleComplete={toggleEventCompletion}
            onUpdate={updateEvent}
            onDelete={deleteEvent}
          />
        </div>
      </div>
    </div>
  )
}
export default App
