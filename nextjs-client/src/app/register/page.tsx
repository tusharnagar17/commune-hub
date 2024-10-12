"use client"
import CommuneHubIcon from '@/components/CommuneHubIcon'
import { registerService } from '@/services/authService'
import { useRouter } from 'next/navigation'
import React, { useState } from 'react'

const page = () => {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null)

  const router = useRouter()

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
        const data = await registerService({username, email, password})
        localStorage.setItem('token', data.token);
        router.push('/dashboard')
    } catch (error: any) {
        setError(error.message)
    } finally {
        setLoading(false);
    }
  }

  return (
    <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
    <div className="sm:mx-auto sm:w-full sm:max-w-sm">
      <CommuneHubIcon />
      <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight">
        Sign Up
      </h2>
    </div>

    <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
      <form action="#" method="POST" className="space-y-6">
        {/* Username */}
        <div>
          <label htmlFor="username" className="block text-md font-medium leading-6">
            Username
          </label>
          <div className="mt-2">
            <input
              id="username"
              name="username"
              type="text"
              onChange={(e)=> {setUsername(e.target.value)}}
              required
              autoComplete="email"
              className="block w-full rounded-md border-0 py-1.5 px-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            />
          </div>
        </div>
        {/* Email */}
        <div>
          <label htmlFor="email" className="block text-md font-medium leading-6">
            Email address
          </label>
          <div className="mt-2">
            <input
              id="email"
              name="email"
              type="email"
              onChange={(e)=> {setEmail(e.target.value)}}
              required
              autoComplete="email"
              className="block w-full rounded-md border-0 py-1.5 px-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            />
          </div>
        </div>
        {/* Password */}
        <div>
          <div className="flex items-center justify-between">
            <label htmlFor="password" className="block text-md font-medium leading-6">
              Password
            </label>
            <div className="text-sm">
              <a href="#" className="font-semibold text-indigo-600 hover:text-indigo-500">
                Forgot password?
              </a>
            </div>
          </div>
          <div className="mt-2">
            <input
              id="password"
              name="password"
              type="password"
              onChange={(e)=> {setPassword(e.target.value)}}
              required
              autoComplete="current-password"
              className="block w-full rounded-md border-0 py-1.5 px-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            />
          </div>
        </div>
        {/* Custom Error */}
        {error && (
            <div className='text-red-500 text-sm'>
                {error}
            </div>
        )}
        <div>
          <button
            type="submit"
            onClick={(e)=> handleRegister(e)}
            className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >
            Sign in
          </button>
        </div>
      </form>

      <p className="mt-10 text-center text-md text-gray-500">
        Already a user?{' '}
        <a href="/login" className="font-semibold leading-6 text-indigo-600 hover:text-indigo-500">
          Login
        </a>
       {" "} Now.
      </p>

    </div>
  </div>
  )
}

export default page