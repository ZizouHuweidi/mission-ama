import prisma from '../db'
import { comparePasswords, createJWT, hashPassword } from '../modules/auth'

export const createEmployee = async (req, res) => {

  const { email, password, name } = req.body

  const employee = await prisma.employee.create({
    data: {
      email: email,
      password: await hashPassword(password),
      name: name,
    }
  })

  const token = createJWT(employee)
  res.json({ token })
}

export const signin = async (req, res) => {
  const employee = await prisma.employee.findUnique({
    where: {
      email: req.body.email
    }
  })

  const isValid = await comparePasswords(req.body.password, employee.password)

  if (!isValid) {
    res.status(401)
    res.json({ message: 'nope' })
    return
  }

  const token = createJWT(employee)
  res.json({ token })
}

// Get all
export const getEmployees = async (req, res) => {
  try {
    const employee = await prisma.employee.findMany()

    res.json(employee)
  } catch (error) {
    res.status(500).json({
      message: "stinky poo poo"
    })
  }
}

// Get one
export const getOneEmployee = async (req, res) => {
  const { id } = req.params
  const member = await prisma.employee.findFirst({
    where: { id: Number(id) },
  })
  res.json({
    member
  })
}


// Update one
export const updateEmployee = async (req, res) => {
  // const { email, role, firstName, lastName, dob, phone, membershipFee } = req.body
  try {
    const employee = await prisma.employee.update({
      where: {
        id: Number(req.params.id),
      },
      data: {
      },
    })
    res.status(200).json(employee)
  } catch (error) {
    res.status(400).json({ msg: error.message })
  }
}

// Delete one
export const deleteEmployee = async (req, res) => {
  const { id } = req.params
  const employee = await prisma.employee.delete({
    where: {
      id: Number(id),
    },
  })
  res.json(employee)
}
