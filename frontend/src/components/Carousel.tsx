import  { useState } from 'react'
import { FaAngleLeft, FaAngleRight } from 'react-icons/fa';
import { Link } from 'react-router-dom';

export default function Carousel() {
    const images: string[] = [
        "https://wallpapersmug.com/download/3840x2160/f12332/books.jpg",
        "https://c1.wallpaperflare.com/preview/127/366/443/library-book-bookshelf-read.jpg",
        "https://png.pngtree.com/background/20231030/original/pngtree-virtual-bookstore-3d-illustration-of-online-shopping-for-books-on-a-picture-image_5787307.jpg",
      ];
    
      const [currentSlide, setCurrentSlide] = useState<number>(0);
    
      const nextSlide = () => {
        setCurrentSlide((prevSlide) =>
          prevSlide === images.length - 1 ? 0 : prevSlide + 1
        );
      };
    
      const prevSlide = () => {
        setCurrentSlide((prevSlide) =>
          prevSlide === 0 ? images.length - 1 : prevSlide - 1
        );
      };
    
      let Carousel = {
        backgroundImage: `linear-gradient(rgba(0, 0, 0, 0.4), rgba(0, 0, 0, 0.4)), url(${images[currentSlide]})`,
        backgroundSize: "cover",
        backgroundPosition: "center",
      };

  return (
    <div className={`max-w-screen bg-background `} style={Carousel}>
        <div className=" sm:w-[640px] md:w-[768px] lg:w-[1024px] xl:w-[1280px] 2xl:w-[1536px] mx-auto">
          <div className="flex flex-col justify-around w-full h-[400px] relative py-8 px-8">
            <div className="flex flex-col items-center gap-4">
              <p className=" text-4xl font-black text-background outline-dotted p-2 shadow-2xl">
                BEST BOOK SELLER
              </p>
              <p className="font-extralight text-background ">
                Save up to 100%
              </p>
            </div>

            <div className="flex justify-center">
              <Link to={"/product"}>

              <button className="bg-secondary px-12 py-2 rounded-lg">
                <p className="font-semibold">Shop Now</p>
              </button>
              </Link>
            </div>

            <div className="p-4 absolute text-5xl text-background h-full flex items-center Z-10 top-1 left-1">
              <FaAngleLeft
                onClick={prevSlide}
                className="transition ease-in-out delay-450 rounded-full p-2 hover:bg-white hover:text-text"
              />
            </div>

            <div className=" p-4 absolute text-5xl right-1 text-background h-full flex items-center Z-10 top-1">
              <FaAngleRight
                onClick={nextSlide}
                className="transition ease-in-out delay-450 rounded-full p-2 hover:bg-white hover:text-text"
              />
            </div>
          </div>
        </div>
      </div>
  )
}
