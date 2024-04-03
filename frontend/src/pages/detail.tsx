import Navbar from "../components/layouts/navbar";
import Footer from "../components/layouts/footer";
import DynamicLayout from "../components/layouts/dynamicLayout";
import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { Book } from "../model/model";
import { apiGetBook } from "../api/getBooks";
import PageNotFound from "../components/404Page";
export default function Detail() {
  let { id } = useParams();
  const [book, setBook] = useState<Book>();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const bookData = await apiGetBook(Number(id));
        setBook(bookData);
      } catch (error: any) {
        setBook(error.message);
      }
    };
    fetchData();
  }, [id]);

  return (
    <div>
      <Navbar />
      <DynamicLayout border={false}>
        {typeof book === "string" ? (
          <PageNotFound />
        ) : (
          book &&
          book.Cover && (
            <div className="flex py-12">
              <div className="min-w-[50%] flex justify-center align-center">
                <img
                  className="object-contain h-96"
                  src={`https://covers.openlibrary.org/b/id/${book.Cover}-L.jpg`}
                  alt="Book Cover"
                />
              </div>
              <div className="min-w-[50%] p-2 flex flex-col gap-4">
                <p className=" text-4xl font-black">{book.Title}</p>
                <p className="w-[70%]">
                  Lorem ipsum dolor sit amet consectetur adipisicing elit. Enim
                  et dolore, dolorum eveniet nisi eligendi itaque soluta nobis
                  facilis ad blanditiis rem nesciunt tenetur minima, eaque
                  possimus distinctio earum est!
                </p>
                <p className="font-bold text-lg">Genre : {book.Genre}</p>
                <p className="text-5xl font-medium ">${book.Price}</p>
                <div></div>
              </div>
            </div>
          )
        )}
      </DynamicLayout>

      <Footer />
    </div>
  );
}
