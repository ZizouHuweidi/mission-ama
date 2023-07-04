import express from 'express'
// import router from './router'
import morgan from 'morgan'
import cors from 'cors'
import { protect } from './modules/auth'
import { createMember, signin } from './controllers/member'
import { deleteMember, getOneMember, getMembers, updateMember } from './controllers/member'
import { createActivity, deleteActivity, getOneActivity, getActivities, updateActivity } from './controllers/activity'
import { takeAttendance, updateAttendance, deleteAttendance, getMemberAttendance, getActivityAttendance, getAllAttendance } from './controllers/attendance'
import { createReceipt, getReceipts, getOneReceipt, updateReceipt, deleteReceipt } from './controllers/receipt'
import { createNotification, deleteNotification, getOneNotification, getNotifications, updateNotification } from './controllers/notification'

const app = express()

app.use(cors())
app.use(morgan('dev'))
app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.get('/', (req, res, next) => {
  setTimeout(() => {
    next(new Error('hello'))
  }, 1)
})

// app.use('/api', protect, router)
app.post('/signin', signin)

/**
 * Member
 */

app.post('/member', createMember)
app.get('/member', getMembers)
app.get('/member/:id', getOneMember)
app.patch('/member/:id', updateMember)
app.delete('/member/:id', deleteMember)

/**
 * Activity
 */

app.post('/activity', createActivity)
app.get('/activity', getActivities)
app.get('/activity/:id', getOneActivity)
app.patch('/activity/:id', updateActivity)
app.delete('/activity/:id', deleteActivity)

/**
 * Attendance
 */

app.post('/attendance', takeAttendance)
app.get('/attendance', getAllAttendance)
app.get('/attendance/:id', getActivityAttendance)
app.get('/attendance/member/:id', getMemberAttendance)
app.patch('/attendance/:id', updateAttendance)
app.delete('/attendance/:id', deleteAttendance)



/**
 * Receipt
 */
app.post('/receipt', createReceipt)
app.get('/receipt', getReceipts)
app.get('/receipt/:id', getOneReceipt)
app.patch('/receipt/:id', updateReceipt)
app.delete('/receipt/:id', deleteReceipt)



/**
 * notifications
 */
// TODO: Fininsh notification implementation
app.post('/notification', createNotification)
app.get('/notification', getNotifications)
app.get('/notification/:id', getOneNotification)
app.patch('/notification/:id', updateNotification)
app.delete('/notification/:id', deleteNotification)


app.use((err, req, res, next) => {
  console.log(err)
  res.json({ message: `had an error: ${err.message}` })
})

export default app
