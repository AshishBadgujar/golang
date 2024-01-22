import Image from 'next/image'
import { Inter } from '@next/font/google'
import Link from 'next/link'
import { useState } from 'react'


export default function Home() {
  const [username, setUsername] = useState('')
  return (
    <>
      <div className='w-full bg-white h-screen flex justify-center items-center flex-col space-y-4'>
        <div><Image src="/go-8.svg" alt="open AI Logo" width={40} height={40} /></div>
        <div className='flex flex-col text-black justify-center items-center space-y-2'>
          <h1 className='text-4xl font-bold'>Welcome to Golang Chat</h1>
          <p>Give us your username befor start</p>
          <br />
          <div className='flex flex-col space-y-4 w-80'>
            <input placeholder='Username' value={username} onChange={(e) => setUsername(e.target.value)} className='border border-gray-400 rounded text-black p-3 bg-white' />
            <Link href={username.length == 0 ? `#` : `/home?username=${username}`} className='bg-teal-600 hover:bg-teal-700 rounded-sm font-light text-white p-4 px-4 text-center' >Continue</Link>
          </div>
        </div>
      </div>
    </>
  )
}
