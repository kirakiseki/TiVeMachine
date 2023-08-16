package impl

import (
	"errors"
	"program/dao"
	"program/model/dto"
	"program/utils"
)

type ProgramServiceImpl struct{}

func (p *ProgramServiceImpl) List() (dto.Response, error) {
	resp := dto.Response{}

	list, err := dao.GetProgramList()
	if err != nil {
		utils.Fail(&resp, 50021, "查询失败")
	}

	utils.Success(&resp)
	listResp := dto.ListResponse{
		List: list,
	}
	resp.Data = listResp
	return resp, nil
}

func (p *ProgramServiceImpl) ListByPeriod(startTime, endTime uint) (dto.Response, error) {
	resp := dto.Response{}

	list, err := dao.GetProgramListByPeriod(startTime, endTime)
	if err != nil {
		utils.Fail(&resp, 50027, "查询失败")
	}

	utils.Success(&resp)
	listResp := dto.ListResponse{
		List: list,
	}
	resp.Data = listResp
	return resp, nil
}

func (p *ProgramServiceImpl) ListByCategory(category string) (dto.Response, error) {
	resp := dto.Response{}

	list, err := dao.GetProgramListByCategory(category)
	if err != nil {
		utils.Fail(&resp, 50028, "查询失败")
	}

	utils.Success(&resp)
	listResp := dto.ListResponse{
		List: list,
	}
	resp.Data = listResp
	return resp, nil
}

func (p *ProgramServiceImpl) ListByChannel(channelID uint) (dto.Response, error) {
	resp := dto.Response{}

	list, err := dao.GetProgramListByChannel(channelID)
	if err != nil {
		utils.Fail(&resp, 50029, "查询失败")
	}

	utils.Success(&resp)
	listResp := dto.ListResponse{
		List: list,
	}
	resp.Data = listResp
	return resp, nil
}

func (p *ProgramServiceImpl) Info(programID uint) (dto.Response, error) {
	resp := dto.Response{}

	if !dao.HasProgramByID(programID) {
		utils.Fail(&resp, 50022, "节目不存在")
		return resp, errors.New("节目不存在")
	}

	program, err := dao.GetProgramInfo(programID)
	if err != nil {
		utils.Fail(&resp, 50023, "获取节目信息失败")
		return resp, err
	}

	utils.Success(&resp)
	infoResp := dto.ProgramInfoResponse{
		Info: program,
	}
	resp.Data = infoResp

	return resp, nil
}

func (p *ProgramServiceImpl) Search(programName string) (dto.Response, error) {
	resp := dto.Response{}

	list, err := dao.SearchProgram(programName)
	if err != nil {
		utils.Fail(&resp, 50024, "查询失败")
		return resp, err
	}

	utils.Success(&resp)
	listResp := dto.ListResponse{
		List: list,
	}
	resp.Data = listResp
	return resp, nil
}

func (p *ProgramServiceImpl) Add(program dto.ProgramDTO) (dto.Response, error) {
	resp := dto.Response{}

	if dao.HasProgramByName(program.Name) {
		utils.Fail(&resp, 50025, "节目已存在")
		return resp, errors.New("节目已存在")
	}

	err := dao.AddProgram(program)
	if err != nil {
		utils.Fail(&resp, 50026, "添加节目失败")
		return resp, err
	}

	utils.Success(&resp)
	return resp, nil
}

func (p *ProgramServiceImpl) Delete(programID uint) (dto.Response, error) {
	resp := dto.Response{}

	if !dao.HasProgramByID(programID) {
		utils.Fail(&resp, 50030, "节目不存在")
		return resp, errors.New("节目不存在")
	}

	err := dao.DeleteProgram(programID)
	if err != nil {
		utils.Fail(&resp, 50031, "删除节目失败")
		return resp, err
	}

	utils.Success(&resp)
	return resp, nil
}
