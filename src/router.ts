
import { Router } from 'express'
import { body, oneOf, validationResult } from "express-validator"
import { deleteMember, getOneMember, getMembers, updateMember } from './controllers/member'
import { createActivity, deleteActivity, getOneActivity, getActivities, updateActivity } from './controllers/activity'
import { takeAttendance, updateAttendance, deleteAttendance, getMemberAttendance, getActivityAttendance, getAllAttendance } from './controllers/attendance'
import { createReceipt, getReceipts, getOneReceipt, updateReceipt, deleteReceipt } from './controllers/receipt'
import { createNotification, deleteNotification, getOneNotification, getNotifications, updateNotification } from './controllers/notification'
import { handleInputErrors } from './modules/middleware'


//TODO: Finish all needed validation and error handeling

const router = Router()

/**
 * Member
 */
router.get('/member', getMembers)
router.get('/member/:id', getOneMember)
router.patch('/member/:id', updateMember)
router.delete('/member/:id', deleteMember)

/**
 * Activity
 */

router.post('/activity', createActivity)
router.get('/activity', getActivities)
router.get('/activity/:id', getOneActivity)
router.patch('/activity/:id', updateActivity)
router.delete('/activity/:id', deleteActivity)

/**
 * Attendance
 */

router.post('/attendance', takeAttendance)
router.get('/attendance', getAllAttendance)
router.get('/attendance/:id', getActivityAttendance)
router.get('/attendance/member/:id', getMemberAttendance)
router.patch('/attendance/:id', updateAttendance)
router.delete('/attendance/:id', deleteAttendance)



/**
 * Receipt
 */
router.post('/receipt', createReceipt)
router.get('/receipt', getReceipts)
router.get('/receipt/:id', getOneReceipt)
router.patch('/receipt/:id', updateReceipt)
router.delete('/receipt/:id', deleteReceipt)



/**
 * notifications
 */
// TODO: Fininsh notification implementation
router.post('/notification', createNotification)
router.get('/notification', getNotifications)
router.get('/notification/:id', getOneNotification)
router.patch('/notification/:id', updateNotification)
router.delete('/notification/:id', deleteNotification)

export default router
