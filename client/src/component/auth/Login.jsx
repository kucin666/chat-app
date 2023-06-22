import { useContext, useState } from "react"
import { useMutation } from "react-query"
import { API, setAuthToken } from "../../config/Api"
import { UserContext } from "../../context/UserContext"
import { alertLoginSucces, alertLoginFailed } from "../alert/Alert"
import { useNavigate } from "react-router-dom"


function Login() {
  const [message, setMessage] = useState(null)
  const [, dispatch] = useContext(UserContext)
  const navigate = useNavigate()

  const [valueLogin, setValueLogin] = useState({
    email: "",
    password: "",
  })

  const { email, password } = valueLogin

  const handleOnChangeLogin = (e) => {
    setValueLogin({
      ...valueLogin,
      [e.target.name]: e.target.value
    })
  }

  function redirectPage() {
    setTimeout(() => {
      navigate("/room")
    }, 3000); // 3000 milidetik = 3 detik
  }

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault()

      const response = await API.post("/login", valueLogin)
      console.log("login success")
      console.table(response.data.data)

      // send data to useContext
      dispatch({
        type: 'LOGIN_SUCCESS',
        payload: response.data.data,
      })

      // set authorization token header
      setAuthToken(localStorage.token)

      setValueLogin({
        email: "",
        password: "",
      });
        
      setMessage(alertLoginSucces);

      redirectPage()

    } catch(err) {
      setMessage(alertLoginFailed);
      console.log("login failed : ", err);
    }
  })

  const token = localStorage.getItem('token')
  console.log(token)

  return (
    <div className="card-body">
      <form onSubmit={(e) => handleSubmit.mutate(e)}>
        <div className="form-control">
          <label className="label">
            <span className="label-text">Email</span>
          </label>
          <input onChange={handleOnChangeLogin} type="email" name="email" value={email} placeholder="email" className="input input-bordered" />
        </div>
        <div className="form-control">
          <label className="label">
            <span className="label-text">Password</span>
          </label>
          <input onChange={handleOnChangeLogin} type="password" name="password" value={password} placeholder="password" className="input input-bordered" />
          <label className="label">
            <a href="#" className="label-text-alt link link-hover">Don&#39;t have an account ? Click Here</a>
          </label>
        </div>
        {message && message}
        <div className="form-control mt-6">
          <button type="submit" className="btn btn-primary">Login</button>
        </div>
      </form>
    </div>
  )
}

export default Login