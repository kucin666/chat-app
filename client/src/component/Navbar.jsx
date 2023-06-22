import { useQuery } from "react-query"
import { API } from "../config/Api"

function Navbar() {

  let { data: profileData } = useQuery('profilCacheNavbar', async() => {
    const res = await API.get('/user-by-login')
    return res.data.data
  })

  console.log(profileData)

  return (
    <div>
      <nav className="">
        {/* <div className=""> */}
          <div className="flex justify-center">
              <div className="m-3 w-[640px] flex justify-end pr-4 items-center">
                <h1 className="mr-3 text-lg">{profileData?.name}</h1>
                <div className="avatar online">
                  <div className="w-12 rounded-full">
                    <img src={profileData?.profile_image} />
                  </div>
                </div>
              </div>
          </div>
        {/* </div> */}
      </nav>
    </div>
  )
}

export default Navbar