import prisma from "../db"

export const createMission = async (req, res) => {
  const { startDate, endDate, destination, purpose, transport, employeeId } = req.body
  try {
    const mission = await prisma.mission.create({
      data: {
        startDate: startDate,
        endDate: endDate,
        destination: destination,
        purpose: purpose,
        transport: transport,
        employeeId: employeeId
      },
    });
    res.status(201).json(mission)
  } catch (error) {
    res.status(400).json({ msg: error.message })
  }
}

export const getMissions = async (req, res) => {
  try {
    const response = await prisma.mission.findMany()
    res.status(200).json(response)
  } catch (error) {
    res.status(500).json({ msg: error.message })
  }
}

export const getOneMission = async (req, res) => {
  try {
    const response = await prisma.mission.findUnique({
      where: {
        id: Number(req.params.id),
      },
    })
    res.status(200).json(response)
  } catch (error) {
    res.status(404).json({ msg: error.message })
  }
}


export const updateMission = async (req, res) => {
  // const { name, description, location, date, type } = req.body
  try {
    const mission = await prisma.mission.update({
      where: {
        id: Number(req.params.id),
      },
      data: {

      },
    })
    res.status(200).json(mission)
  } catch (error) {
    res.status(400).json({ msg: error.message })
  }
}

export const deleteMission = async (req, res) => {
  try {
    const mission = await prisma.mission.delete({
      where: {
        id: Number(req.params.id),
      },
    })
    res.status(200).json(mission)
  } catch (error) {
    res.status(400).json({ msg: error.message })
  }
}
