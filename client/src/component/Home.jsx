import { useState } from "react"
import Login from "./auth/Login"
import Register from "./auth/Register"

function Home() {
  const [showRegister] = useState(false)

  return (
    <div className="flex justify-center bg-base-200">
      <div className="w-[800px]">
        <div className="hero min-h-screen bg-base-200">
          <div className="hero-content flex-col lg:flex-row-reverse">
            <div className="text-center lg:text-left">
              <h1 className="text-5xl font-bold">Hi Chat!</h1>
              <p className="py-6">an online communication platform that helps you connect with important people in your life easily and securely. With advanced features and an intuitive interface, we ensure that your chat experience becomes more enjoyable and efficient.</p>
            </div>
            <div className="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
              {showRegister ? <Register /> : "" }
              <Login />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Home