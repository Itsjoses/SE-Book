import axios from "axios";
import { useEffect, useState } from "react";

interface Book {
    title: string,
    coverId: string,
    genre : number,
    price: number
}

export default function GetJson() {

    const [books, setBooks] = useState<Book[]>([])
 
    useEffect(() => {
        axios.get(`https://openlibrary.org/search.json?title=adventure&limit=300`)
            .then((res:any) => {
                console.log(res);
                const fetchData = async () => {
                    const bookPromises = res.data.docs.map(async (data:any) => {
                        if (data.cover_i != null){
                            const bookData = {
                                Title: data.title,
                                Cover: data.cover_i,
                                GenreID: Math.floor(Math.random() * 4) + 1,
                                Price: Math.floor(Math.random() * (200 - 50 + 1)) + 50
                            };
                            
                            return bookData;
                        }
                    });

                    const booksData = await Promise.all(bookPromises);
                    setBooks(booksData);
                };

                fetchData();
            })
            .catch((error: any) => {
                console.error("Error fetching books:", error);
            });
    }, []);

    const checkBook = () => {
        console.log(JSON.stringify(books));
    };

  return <div >
    <button onClick={checkBook}>book</button>
  </div>;
}
