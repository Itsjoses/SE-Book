import React from 'react'

interface children {
  border: boolean,
  children: React.ReactNode
}

export default function DynamicLayout({border,children}: children) {
  return (
    <div className={`max-w-screen  ${border ? `border-b-2 border-[#E6EAF3]` : ``} `}>
      <div className='sm:w-[640px] md:w-[768px] lg:w-[1024px] xl:w-[1280px] 2xl:w-[1536px] mx-auto'>
        {children}
      </div>
    </div>
  )
}
