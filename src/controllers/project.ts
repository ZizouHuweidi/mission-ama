import prisma from "../db"

export const createProject = async (req, res) => {
  const { name, description } = req.body
  try {
    const project = await prisma.project.create({
      data: {
        name: name,
        description: description
      },
    });
    res.status(201).json(project)
  } catch (error) {
    res.status(400).json({ msg: error.message })
  }
}

export const getProjects = async (req, res) => {
  try {
    const response = await prisma.project.findMany()
    res.status(200).json(response)
  } catch (error) {
    res.status(500).json({ msg: error.message })
  }
}

export const getOneProject = async (req, res) => {
  try {
    const response = await prisma.project.findUnique({
      where: {
        id: Number(req.params.id),
      },
    })
    res.status(200).json(response)
  } catch (error) {
    res.status(404).json({ msg: error.message })
  }
}


export const updateProject = async (req, res) => {
  const { name, description } = req.body
  try {
    const project = await prisma.project.update({
      where: {
        id: Number(req.params.id),
      },
      data: {
        name: name,
        description: description
      },
    })
    res.status(200).json(project)
  } catch (error) {
    res.status(400).json({ msg: error.message })
  }
}

export const deleteProject = async (req, res) => {
  try {
    const project = await prisma.project.delete({
      where: {
        id: Number(req.params.id),
      },
    })
    res.status(200).json(project)
  } catch (error) {
    res.status(400).json({ msg: error.message })
  }
}
