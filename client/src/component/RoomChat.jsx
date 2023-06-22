import { useMutation, useQuery } from "react-query"
import { Link, useParams } from "react-router-dom"
import { API, setAuthToken } from "../config/Api"
import { useEffect, useState } from "react"
import { alertSendingMessageFailed, alertSendingMessageSuccess } from "./alert/Alert"

function RoomChat() {
  const { id } = useParams()
  const [message, setMessage] = useState(null)
  const [valueMessage, setValueMessage] = useState({
    title: "",
  })

  const { title } = valueMessage 

  const handleOnChangeMessage = (e) => {
    setValueMessage({
      ...valueMessage,
      [e.target.name]: e.target.value
    })
  }

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault()

      const config = {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('token')}`,
        },
      }

      // store data with FormData as object
      const formData = new FormData()
      formData.set('title', valueMessage.title)
      formData.set('room_id', Number(id))

      const response = await API.post("/chat", formData, config)
      console.log("sending message success")
      console.table(response.data.data)

      setValueMessage({
        title: "",
      });
        
      setMessage(alertSendingMessageSuccess);

    } catch(err) {
      setMessage(alertSendingMessageFailed);
      console.log("sending message failed : ", err);
    }
  })

  let { data: chatlist, refetch } = useQuery('chatChache', async() => {
    const res = await API.get(`/room/${id}/chats`)
    console.table(res.data.data)
    return res.data.data
  }) 

  let { data: userByLogin } = useQuery('userByLoginCache', async() => {
    const res = await API.get('/user-by-login', setAuthToken(localStorage.token))
    console.table(res.data.data)
    return res.data.data
  })

  useEffect(() => {
    refetch()
  }, [valueMessage])

  return (
    <div className="flex justify-center bg-base-100">
      <div className="w-full md:w-[800px] bg-base-200">
        <div className="md:mx-5">
          <div className="flex items-center justify-between md:justify-normal">
            <div className="w-[30%] md:w-[10%]">
              <Link to={`http://localhost:5173/room`}>
                <button className="btn btn-sm">Back</button>
              </Link>
            </div>
            <h1 className="md:text-3xl font-bold text-center w-[40%] md:w-[80%]">Room Chat</h1>
            <div className="w-[30%] md:w-[10%]">
              <div className="flex justify-center">
                <div className="md:m-3 md:w-[640px] flex md:justify-end md:pr-4 items-center">
                  <h1 className="md:mr-3 md:text-lg">Thxrhmn</h1>
                  <div className="avatar online">
                    <div className="w-8 rounded-full">
                      <img src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg" />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div className="flex flex-col justify-between h-[90%]">
            <div className="m-3 md:m-0">
              {chatlist?.map((item) => (
                <div className="" key={item?.id}>
                  {item?.user_id !== userByLogin?.id ? 
                    (
                      <div key={item?.id} className="chat chat-start">
                        <div className="chat-image avatar">
                          <div className="w-10 rounded-full">
                            <img src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg" />
                          </div>
                        </div>
                        <div className="chat-header">
                          {item?.user?.name}
                          <time className="text-xs opacity-50 ml-1">12:45</time>
                        </div>
                        <div className="chat-bubble">{item?.title}</div>
                      </div>
                    ): 
                    (
                      <div key={item?.id} className="chat chat-end">
                        <div className="chat-image avatar">
                          <div className="w-10 rounded-full">
                            <img src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg" />
                          </div>
                        </div>
                        <div className="chat-header">
                          {item?.user?.name}
                          <time className="text-xs opacity-50 ml-1">12:46</time>
                        </div>
                        <div className="chat-bubble">{item?.title}</div>
                      </div>
                    )
                  }
                </div>
              ))}
            </div>
              

            <div className="flex justify-center my-10">
              <div className="w-full m-3 md:m-0 md:w-[400px]">
                <form onSubmit={(e) => handleSubmit.mutate(e)} className="flex flex-col">
                  <textarea onChange={handleOnChangeMessage} name="title" value={title} className="textarea textarea-info mb-3" placeholder="Write message here"></textarea>
                  {message && message}
                  <button type="submit" className="btn btn-active btn-neutral mt-2">Send Message!</button>
                </form>
              </div>
            </div>

          </div>

        </div>
      </div>
    </div>
  )
}

export default RoomChat