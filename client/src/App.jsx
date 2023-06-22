import { Routes, Route, useNavigate } from "react-router-dom"
import { useEffect, useState, useContext } from "react"

import { UserContext } from "./context/UserContext"
import { API, setAuthToken } from "./config/Api"

import Home from './component/Home'
import Room from "./component/Room"
import { PrivateRouteLogin } from "./component/auth/PrivateRoute"
import RoomChat from "./component/RoomChat"

function App() {
  const [state, dispatch] = useContext(UserContext);
  const [isLoading, setIsLoading] = useState(true)

  let navigate = useNavigate()

  useEffect(() =>{
    // Redirect Auth but just when isLoading is false
    if (!isLoading) {
      if (state.isLogin === false) {
        navigate("/")
      }
    }
  }, [isLoading])

  useEffect(() => {
    if (localStorage.token) {
      setAuthToken(localStorage.token)
      checkUser()
    } else {
      setIsLoading(false)
    }
  }, [])

  const checkUser = async () => {
    try {
      const response = await API.get('/check-auth')
      console.log("check user succes", response)
      // get user data
      let payload = response.data.data
      // get token from localstorage
      payload.token = localStorage.token
      // send data to user context
      dispatch({
        type: 'USER_SUCCESS',
        payload,
      })
      setIsLoading(false)
    } catch(error) {
      console.log("check user failed: ", error)
      dispatch({
        type: 'AUTH_ERROR',
      })
      setIsLoading(false)
    }
  }


  return (
    <>
      {isLoading ? null : (
        <Routes>
          <Route path="/" element={<Home />} />  

          <Route element={<PrivateRouteLogin />} >
            <Route path="/room" element={<Room />} />
            <Route path="/room/:id" element={<RoomChat />} />
          </Route>

        </Routes>
      )}
    </>
  )
}

export default App
