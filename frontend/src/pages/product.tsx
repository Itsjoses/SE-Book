import Footer from "../components/layouts/footer";
import Navbar from "../components/layouts/navbar";
import DynamicLayout from "../components/layouts/dynamicLayout";
import Card from "../components/card";
import { useState, useEffect } from "react";
import { Book } from "../model/model";
import { apiSearchBook } from "../api/getBooks";
import { IoSearch } from "react-icons/io5";
import PageNotFound from "../components/404Page";

export default function Product() {
  const [allBooks, setAllBooks] = useState<Book[] | string>([]);
  const [search,setSearch] = useState<string>("") 
  useEffect(() => {
    const fetchData = async () => {
      try{

        const booksData = await apiSearchBook(search);
        setAllBooks(booksData);
      } catch(error: any){
        setAllBooks(error.message);
      }
    };
    fetchData();
  }, [search]);

  return (
    <div>
      <Navbar />
      <DynamicLayout border={false}>
        <div className="flex justify-center my-4">
          <div className="p-2.5 w-[50rem] bg-secondary/30 flex items-center gap-2 rounded-lg justify-center">
            <IoSearch className="text-2xl" />
            <input
              type="search"
              className="bg-transparent w-full outline-none placeholder-black"
              placeholder="Find Your Book..."
              onChange={(e) => {setSearch(e.target.value)}}
            ></input>
          </div>
        </div>
        <div className="py-4 px-8 flex flex-col gap-4">
          <div className="font-bold text-xl italic">Book Products</div>
          {typeof allBooks === "string" ? (
              <PageNotFound/>
            ) : (
              <div className="grid gap-8 lg:grid-cols-6 md:grid-cols-4 sm:grid-cols-2 grid-cols-1">
                {

                  allBooks.map((book: Book,key) => (
                    <Card book={book} key={key}/>
                    ))
                  }
            </div>
            )}
        </div>
      </DynamicLayout>
      <Footer />
    </div>
  );
}
