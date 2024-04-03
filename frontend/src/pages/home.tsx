import DynamicLayout from "../components/layouts/dynamicLayout";
import Footer from "../components/layouts/footer";
import Navbar from "../components/layouts/navbar";
import Card from "../components/card";
import CardGenre from "../components/cardGenre";
import Carousel from "../components/Carousel";
import { useEffect, useState } from "react";
import { Book } from "../model/model";
import { apiGetBooks } from "../api/getBooks";
import { BiGame } from "react-icons/bi";
import { GiAxeSword } from "react-icons/gi";
import { BiGhost } from "react-icons/bi";
import { GiBloodyStash } from "react-icons/gi";
import PageNotFound from "../components/404Page";
export default function Home() {
  const [allBooks, setAllBooks] = useState<Book[] | string>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const booksData = await apiGetBooks();
        setAllBooks(booksData);
      } catch (error: any) {
        setAllBooks(error.message);
      }
    };
    fetchData();
  }, []);

  return (
    <div>
      <Navbar />
      <Carousel />
      <DynamicLayout border={false}>
        <div className="py-4 px-8 flex flex-col gap-4">
          <div className="font-bold text-xl italic">Genre {`>`} </div>
          <div className="grid grid-cols-4 gap-8 ">
            <CardGenre genre={{icon: BiGame, title: "Comedy"}}/>
            <CardGenre genre={{icon: GiAxeSword, title: "Adventure"}}/>
            <CardGenre genre={{icon: BiGhost, title: "Horror"}}/>
            <CardGenre genre={{icon: GiBloodyStash, title: "Thriller"}}/>
          </div>
        </div>
      </DynamicLayout>

      <DynamicLayout border={false}>
        <div className="py-4 px-8 flex flex-col gap-4">
          <div className="font-bold text-xl italic">Book Products</div>
            {typeof allBooks === "string" ? (
              <PageNotFound/>
            ) : (
              <div className="grid gap-8 lg:grid-cols-6 md:grid-cols-4 sm:grid-cols-2 grid-cols-1">
                {

                  allBooks.map((book: Book, key) => (
                    <Card book={book} key={key} />
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