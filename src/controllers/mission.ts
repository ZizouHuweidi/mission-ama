import prisma from "../db"

export const createActivity = async (req, res) => {
  const { name, description, location, date, type } = req.body
  try {
    const activity = await prisma.activity.create({
      data: {
        name: name,
        description: description,
        location: location,
        date: date,
        type: type,
      },
    });
    res.status(201).json(activity)
  } catch (error) {
    res.status(400).json({ msg: error.message })
  }
}

export const getActivities = async (req, res) => {
  try {
    const response = await prisma.activity.findMany()
    res.status(200).json(response)
  } catch (error) {
    res.status(500).json({ msg: error.message })
  }
}

export const getOneActivity = async (req, res) => {
  try {
    const response = await prisma.activity.findUnique({
      where: {
        id: Number(req.params.id),
      },
    })
    res.status(200).json(response)
  } catch (error) {
    res.status(404).json({ msg: error.message })
  }
}


export const updateActivity = async (req, res) => {
  const { name, description, location, date, type } = req.body
  try {
    const activity = await prisma.activity.update({
      where: {
        id: Number(req.params.id),
      },
      data: {
        name: name,
        description: description,
        location: location,
        date: date,
        type: type
      },
    })
    res.status(200).json(activity)
  } catch (error) {
    res.status(400).json({ msg: error.message })
  }
}

export const deleteActivity = async (req, res) => {
  try {
    const activity = await prisma.activity.delete({
      where: {
        id: Number(req.params.id),
      },
    })
    res.status(200).json(activity)
  } catch (error) {
    res.status(400).json({ msg: error.message })
  }
}
