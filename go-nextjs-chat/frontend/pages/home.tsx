import React, { FormEvent, use, useEffect, useRef, useState } from 'react'
import { PaperAirplaneIcon, } from "@heroicons/react/24/outline"
import { useRouter } from 'next/router';
import Image from 'next/image';

interface Message {
    username: string
    text: string
}

const Home = () => {
    const router = useRouter();
    const { username } = router.query;
    const [messages, setMessages] = useState<Message[]>([]);
    const [text, setText] = useState<string>('')
    const socketRef = useRef<WebSocket | null>(null);

    useEffect(() => {
        // Connect to the WebSocket server
        if (!username) return
        socketRef.current = new WebSocket(`ws://localhost:8000/ws?username=${username}`);

        // Listen for incoming messages
        socketRef.current.addEventListener('message', (event) => {
            const message = JSON.parse(event.data);
            console.log("message=", message)
            setMessages((prevMessages) => [...prevMessages, message]);
        });

        // Cleanup the WebSocket connection when the component unmounts
        return () => {
            if (socketRef.current) {
                socketRef.current.close();
            }
        };
    }, [username]);

    useEffect(() => {
        // Scroll to the bottom when messages change
        scrollToBottom();
    }, [messages]);

    const scrollToBottom = () => {
        const scrollContainer = document.getElementById('scrollContainer');
        if (scrollContainer) {
            scrollContainer.scrollTop = scrollContainer.scrollHeight;
        }
    };

    const sendMessage = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (socketRef.current && text.trim() !== '') {
            const message = { username: username, text: text };
            setMessages((prevMessages) => [...prevMessages, message]);
            socketRef.current.send(JSON.stringify(message));
            setText('')
        }
    };


    return (
        <div className='h-screen bg-white text-black flex'>
            <div className='relative flex flex-1 flex-col h-full'>
                <div id='scrollContainer' className="h-full overflow-y-auto pb-24 pt-4">
                    {messages.length == 0 ?
                        <div className='flex flex-row justify-center items-center absolute inset-x-0 top-0 bottom-0'>
                            <Image src="/next.svg" alt="open AI Logo" width={200} height={200} />
                            <h1 className='text-4xl ml-6'>+</h1>
                            <Image src="/go.svg" alt="open AI Logo" width={200} height={200} />
                        </div>
                        :

                        messages.map((msg: Message, index) => (
                            <div className='flex flex-col bg-white text-black' key={index}>
                                {msg.username == "Server" ?
                                    <p className='text-center text-sm text-gray-500'>{msg.text}</p>
                                    :
                                    msg.username == username ?
                                        <div className='w-full flex items-center justify-center bg-gray-200'>
                                            <div className='flex space-x-4 items-center justify-between px-6 py-6 w-1/2'>
                                                <div className='flex space-x-4'>
                                                    <div className='h-8 w-8 bg-teal-600 text-center p-1 px-2 rounded-full text-white'>{msg.username[0]}</div>
                                                    <p>{msg.text}</p>
                                                </div>
                                            </div>
                                        </div>
                                        :
                                        <div className='w-full flex items-center justify-center'>
                                            <div className='flex space-x-4 items-center justify-between px-6 py-6 w-1/2'>
                                                <div className='flex space-x-4'>
                                                    <div className='h-8 w-8 bg-indigo-500 text-center p-1 px-2 rounded-full text-white'>{msg.username[0]}</div>
                                                    <p>{msg.text}</p>
                                                </div>
                                            </div>
                                        </div>
                                }
                            </div>
                        ))

                    }
                </div>
                <div className='absolute bottom-0 inset-x-0 mx-auto px-4 py-6 max-w-3xl'>
                    <form onSubmit={sendMessage} className='bg-white text-black border border-gray-400 rounded flex justify-center items-center space-x-2 shadow-md px-2'>
                        <input placeholder='Message...' className='flex-1 bg-white p-3 border-0 focus:outline-none' value={text} onChange={(e) => setText(e.target.value)} />
                        <PaperAirplaneIcon className='h-4 w-4 text-right -rotate-45' type='submit' />
                    </form>
                </div>
            </div>
        </div>
    )
}

export default Home