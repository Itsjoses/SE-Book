import { IconType } from "react-icons";

interface Genre{
  icon: IconType,
  title: string
}

export default function CardGenre({genre} : {genre: Genre}) {
  return (
    <div className="flex flex-col gap-6 bg-secondary p-8 rounded-lg">
      <div className="h-auto w-auto flex justify-center align-center text-8xl">
        <genre.icon/>
      </div>
      <div className="text-2xl text-accent text-center font-bold  ">{genre.title}</div>
      <div className="text-xs">
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Delectus,
        tenetur recusandae blanditiis error, est consequuntur sint repellat
        ratione ipsa, rerum vel eos molestias cupiditate. Ipsa deleniti ut
        reiciendis odit at?
      </div>
      <div className="flex justify-center">

      <button className="bg-accent/50 flex w-fit px-4 py-2 rounded-xl text-background">Detail</button>
      </div>
    </div>
  );
}
