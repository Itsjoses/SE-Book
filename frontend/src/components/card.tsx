import { useEffect, useState } from "react";
import { Book } from "../model/model";
import axios from "axios";
import { Link } from "react-router-dom";

export default function Card({ book }: { book: Book }) {
  const [coverUrl, setCoverUrl] = useState<string | undefined>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchCover = async () => {
      try {
        const response = await axios.get(
          `https://covers.openlibrary.org/b/id/${book.Cover}-L.jpg`,
          {
            responseType: "blob", 
          }
        );
        const blob = new Blob([response.data], { type: "image/jpeg" });
        setCoverUrl(URL.createObjectURL(blob)); 
        setLoading(false);
      } catch (error) {
        console.error("Error fetching cover:", error);
      }
    };

    fetchCover();

    return () => {
      if (coverUrl) {
        URL.revokeObjectURL(coverUrl);
      }
    };
  }, [book.Cover]);

  return (
    <Link to={`/detail/${book.Id}`}>
      <div className="flex flex-col items-start hover:outline-primary outline-none h-full shadow-lg p-4 bg-background">
        <div className="lg:h-60 md:h-50 sm:h-40 w-full flex justify-center items-center">
          {loading ? (
            <div className="animate-spin"></div>
          ) : (
            <img
              src={coverUrl}
              alt=""
              className="h-full w-full object-contain"
            />
          )}
        </div>
        <div className="py-4 flex flex-col justify-between gap-2">
          <p className="line-clamp-2 font-light text-sm">{book.Title}</p>
          <p className="font-black text-xl">${book.Price}</p>
        </div>
      </div>
    </Link>
  );
}
