import jwt from 'jsonwebtoken'
import bcrypt from 'bcrypt'

export const comparePasswords = (password, hash) => {
  return bcrypt.compare(password, hash)
}

export const hashPassword = (password) => {
  return bcrypt.hash(password, 5)
}

export const createJWT = (member) => {
  const token = jwt.sign({
    id: member.id,
    email: member.email
  },
    process.env.JWT_SECRET
  )
  return token
}

export const protect = (req, res, next) => {
  const bearer = req.headers.authorization

  if (!bearer) {
    res.status(401)
    res.json({ message: 'not authorized' })
    return
  }

  const [, token] = bearer.split(' ')

  if (!token) {
    res.status(401)
    res.json({ message: 'not valid token' })
    return
  }

  try {
    const member = jwt.verify(token, process.env.JWT_SECRET)
    req.member = member
    next()
  } catch (e) {
    console.error(e)
    res.status(401)
    res.json({ message: 'not valid token' })
    return
  }
}
