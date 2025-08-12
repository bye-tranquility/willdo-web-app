import logo from '../assets/logo.svg'

const Header = () => {
  return (
    <header className="header">
      <div className="header-content">
        <div className="header-logo">
          <img src={logo} alt="WillDo! Logo" className="logo-icon" />
          <h1>WillDo!</h1>
        </div>
        <p>Your personal task management app</p>
      </div>
    </header>
  )
}

export default Header
