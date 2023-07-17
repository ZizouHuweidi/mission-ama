
import { Router } from 'express'
import { body, oneOf, validationResult } from "express-validator"
import { createMission, getMissions, getOneMission, updateMission, deleteMission } from './controllers/mission'
import { getEmployees, getOneEmployee, createEmployee, updateEmployee, deleteEmployee } from './controllers/employee'
import { getProjects, createProject, deleteProject, getOneProject, updateProject } from './controllers/project'
import { handleInputErrors } from './modules/middleware'


//TODO: Finish all needed validation and error handeling

const router = Router()

/**
 * Mission
 */
router.get('/mission', getMissions)
router.get('/mission/:id', getOneMission)
router.post('/mission', createMission)
router.patch('/mission/:id', updateMission)
router.delete('/mission/:id', deleteMission)

/**
 * Employee
 */

router.get('/employee', getEmployees)
router.get('/employee/:id', getOneEmployee)
// router.get('/auth/employee/:id', logout)
router.patch('/employee/:id', updateEmployee)
router.delete('/employee/:id', deleteEmployee)

/**
 * Project
 */

router.post('/project', createProject)
router.get('/project', getProjects)
router.get('/project/:id', getOneProject)
router.patch('/project/:id', updateProject)
router.delete('/project/:id', deleteProject)


export default router
