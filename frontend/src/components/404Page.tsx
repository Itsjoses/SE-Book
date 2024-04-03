import { TbError404Off } from 'react-icons/tb'

export default function PageNotFound() {
  return (
    <div className="flex flex-col items-center gap-2 h-96 justify-center text-primary">
                <TbError404Off className="text-7xl"/>  
                <p className="text-2xl font-black ">Content Not Found !</p>
                <p className="text-3xl font-black">Oops ! Network Error Please Contact Your Developer</p>
              </div>
  )
}
