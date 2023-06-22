function RoomChat() {
  return (
    <div className="flex justify-center bg-base-100 h-screen">
      <div className="w-[800px] bg-base-200">
        <div className="mx-5 h-screen">
          <h1 className="text-3xl font-bold text-center mb-5">Room Chat</h1>
          <div className="flex flex-col justify-between h-[90%]">
            <div>
              <div className="chat chat-start">
                <div className="chat-image avatar">
                  <div className="w-10 rounded-full">
                    <img src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg" />
                  </div>
                </div>
                <div className="chat-header">
                  Smith
                  <time className="text-xs opacity-50 ml-1">12:45</time>
                </div>
                <div className="chat-bubble">You were the Chosen One!</div>
              </div>
              <div className="chat chat-end">
                <div className="chat-image avatar">
                  <div className="w-10 rounded-full">
                    <img src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg" />
                  </div>
                </div>
                <div className="chat-header">
                  Jhon
                  <time className="text-xs opacity-50 ml-1">12:46</time>
                </div>
                <div className="chat-bubble">I hate you!</div>
              </div>
            </div>

            <div className="flex justify-center">
              <div className="flex flex-col w-[400px]">
                <textarea className="textarea textarea-info" placeholder="Text here"></textarea>
                <button className="btn btn-active btn-neutral mt-2">Send Message!</button>
              </div>
            </div>

          </div>

        </div>
      </div>
    </div>
  )
}

export default RoomChat