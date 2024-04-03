import React from "react";
import DynamicLayout from "./dynamicLayout";
import { FaYoutube } from "react-icons/fa";
import { FaInstagram } from "react-icons/fa6";
import { FaFacebook } from "react-icons/fa";
export default function Footer() {
  return (
    <DynamicLayout border={false}>
      <div className=" border-t border-accent p-8 flex gap-8 text-text w-full justify-center">
        <div className="font-black italic w-1/3 text-lg text-primary">SEBook</div>
        <div className="flex flex-col gap-1  w-1/3">
          <p className="text-xs text-accent">LET'S CONNECT</p>
          <p className="flex items-center gap-2">
            <FaYoutube className="text-2xl text-primary" /> Youtube
          </p>
          <p className="flex items-center gap-2">
            <FaInstagram className="text-2xl text-primary" /> Instagram
          </p>
          <p className="flex items-center gap-2">
            <FaFacebook className="text-2xl text-primary" /> Facebook
          </p>
        </div>

        <div className="flex flex-col gap-2  w-1/3">
          <p className="text-xs font-light">
            Made with <span className="text-primary">❤</span> by @itsjose.s
          </p>
          <p className="text-xs font-light">
            Copyright © 2024 All Rights Reserved
          </p>
        </div>
      </div>
    </DynamicLayout>
  );
}
