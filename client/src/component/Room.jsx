import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { API } from "../config/Api";
import { useMutation, useQuery } from "react-query";
import { alertCreateRoomFailed, alertCreateRoomSuccess } from "./alert/Alert";

function Room() {
  document.title = "Room List"
  const [message, setMessage] = useState(null)
  const [valueRoom, setValueRoom] = useState({
    name: "",
  })

  const { name } = valueRoom

  const handleOnChangeRoom = (e) => {
    setValueRoom({
      ...valueRoom,
      [e.target.name]: e.target.value
    })
  }

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault()

      // configuration
      const config = {
        headers: {
          'Content-type': 'multipart/form-data',
          'Authorization': `Bearer ${localStorage.getItem('token')}`,
        },
      }

      const response = await API.post("/room", valueRoom, config)
      console.log("create room success")
      console.table(response.data.data)

      setValueRoom({
        name: "",
      });
        
      setMessage(alertCreateRoomSuccess);

    } catch(err) {
      setMessage(alertCreateRoomFailed);
      console.log("create room failed : ", err);
    }
  })

  let {data: roomList, refetch} = useQuery('roomChache', async() => {
    const res = await API.get('/rooms')
    console.log(res.data.data)
    return res.data.data
  })

  useEffect(() => {
    refetch()
  }, [valueRoom])

  return (
    <div className="md:flex md:justify-center bg-base-100 md:h-screen">
      <div className="w-full md:w-[800px] bg-base-200">
        <div className="flex justify-center">
          <div className="m-3 w-[640px] flex justify-end pr-4 items-center">
            <h1 className="mr-3 text-lg">Thxrhmn</h1>
            <div className="avatar online">
              <div className="w-12 rounded-full">
                <img src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg" />
              </div>
            </div>
          </div>
        </div>

        <div className="md:m-5">
          <h1 className="text-3xl text-center font-bold text-white">Room List</h1>
        </div>

        <div className="flex justify-center">
          <div className="md:m-3 md:w-[640px] flex justify-end md:pr-4">
            <button className="btn" onClick={()=>window.my_modal_2.showModal()}>Create Room +</button>
          </div>
        </div>

        <dialog id="my_modal_2" className="modal">
          <form method="dialog" onSubmit={(e) => handleSubmit.mutate(e)} className="modal-box"> 
            <div className="form-control mb-3">
              <label className="label">
                <span className="label-text">Name</span>
              </label>
              <input onChange={handleOnChangeRoom} type="text" name="name" value={name} placeholder="Name" className="input input-bordered" />
            </div>
              {message && message}
            <div className="form-control mt-6">
              <button type="submit" className="btn btn-primary">Create</button>
            </div>
          </form>
          <form method="dialog" className="modal-backdrop">
            <button>close</button>
          </form>
        </dialog>
        
        <div className="flex justify-center">
          <div className="w-full flex flex-col md:flex-row md:flex-wrap gap-5 md:w-[640px]">
            {roomList?.map((item) => (
              <Link to={`${item.id}`} key={item.id}>
                <div className="mx-3 md:m-0 md:w-[200px] bg-gray-800 h-[150px] rounded-md flex justify-center items-center outline-slate-400 outline-2 outline hover:scale-105 cursor-pointer">
                  <h1 className="text-2xl text-white font-semibold">{item?.name}</h1>
                </div>
              </Link>
            ))}
          </div>
        </div>

      </div>
    </div>
  );
}

export default Room;
