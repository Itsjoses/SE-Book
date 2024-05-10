import DynamicLayout from "./dynamicLayout";
import { Link } from "react-router-dom";

export default function Navbar() {
  return (
    <DynamicLayout border={true}>
      <div className="flex flex-col gap-4 p-4">
        <div className="flex justify-around items-center text-text">
          <div className="font-black text-2xl italic">
            <div className=" italic text-primary ">SEBook</div>
          </div>
          <div className="flex font-semibold justify-center gap-12 p-4  ">
            <Link to={"/"}>
              Home
            </Link>
            <Link to={"/product"} className="">
              Product
            </Link>
            <Link to={"/about-us"} className="">
              About Us
            </Link>
            
          </div>
    
          <div className="gap-8 flex items-center font-bold">
            <Link to={"/"} className=" text-primary ">
              Login
            </Link>
            <Link to={"/"} className=" ">
              Sign Up
            </Link>
          </div>
        </div>
      </div>
    </DynamicLayout>
  );
}
