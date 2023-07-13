import prisma from "../db"

export const createMission = async (req, res) => {
  const { startDate, endDate, destination, purpose, transport, employeeId, fuel, external, visa, vaccination, tests, timbre, other_costs, projectId, participants } = req.body
  try {
    const mission = await prisma.mission.create({
      data: {
        employeeId: employeeId,
        projectId: projectId,
        startDate: startDate,
        endDate: endDate,
        participants: participants,
        destination: destination,
        purpose: purpose,
        transport: transport,
        fuel: fuel,
        external: external,
        visa: visa,
        vaccination: vaccination,
        tests: tests,
        timbre: timbre,
        other_costs: other_costs
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
  const { startDate, endDate, destination, purpose, transport, employeeId, fuel, external, visa, vaccination, tests, timbre, other_costs, projectId, participants } = req.body
  try {
    const mission = await prisma.mission.update({
      where: {
        id: Number(req.params.id),
      },
      data: {
        employeeId: employeeId,
        projectId: projectId,
        startDate: startDate,
        endDate: endDate,
        participants: participants,
        destination: destination,
        purpose: purpose,
        transport: transport,
        fuel: fuel,
        external: external,
        visa: visa,
        vaccination: vaccination,
        tests: tests,
        timbre: timbre,
        other_costs: other_costs
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
