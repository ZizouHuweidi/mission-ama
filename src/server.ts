import express from 'express'
import router from './router'
import morgan from 'morgan'
import cors from 'cors'
import { protect } from './modules/auth'
import { createEmployee, signin } from './controllers/employee'

const app = express()

const corsOption = {
  origin: ['http://localhost:5173'],
};

app.use(cors(corsOption))
app.use(morgan('dev'))
app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.get('/', (req, res, next) => {
  setTimeout(() => {
    next(new Error('hello'))
  }, 1)
})

app.use('/api', protect, router)
app.post('/employee', createEmployee)
app.post('/signin', signin)

app.use((err, req, res, next) => {
  console.log(err)
  res.json({ message: `had an error: ${err.message}` })
})


export default app
