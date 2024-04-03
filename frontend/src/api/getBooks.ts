import axios from "axios";
import { Book } from "../model/model";

export async function apiGetBooks() {
  try {
    const response = await axios.get(`${import.meta.env.VITE_API_URL}/getbooks`, {
      headers: {
        Accept: "application/json",
      },
    });

    const bookData: Book[] = response.data.map((data: any) => ({
      Id: data.ID,
      Title: data.Title,
      Cover: data.Cover,
      Price: data.Price,
      Genre: data.Genre.Name,
    }));

    return bookData;
  } catch (error) {
    throw error;
  }
}

export async function apiGetBook(productId: number) {
  try {
    const response: any = await axios.get(
      `${import.meta.env.VITE_API_URL}/getbook/${productId}`,
      {
        headers: {
          Accept: "application/json",
        },
      }
    );
    const bookData: Book = {
      Id: response.data.ID,
      Title: response.data.Title,
      Cover: response.data.Cover,
      Price: response.data.Price,
      Genre: response.data.Genre.Name,
    };
    return bookData;
  } catch (error) {
    throw error;
  }
}

export async function apiSearchBook(searchTerm: string) {
  return await axios
    .post(
      `${import.meta.env.VITE_API_URL}/searchbook`,
      {
        Search: searchTerm,
      },
      {
        headers: {
          Accept: "application/json",
        },
      }
    )
    .then((response) => {
      const bookData: Book[] = response.data.map((data: any) => ({
        Id: data.ID,
        Title: data.Title,
        Cover: data.Cover,
        Price: data.Price,
        Genre: data.Genre.Name,
      }));
      return bookData;
    })
    .catch((error) => {
      throw error;
    });
}
