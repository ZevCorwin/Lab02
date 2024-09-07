package main

import (
	"fmt"
	"strconv"
)

type SinhVien struct {
	Ten      string
	Tuoi     int8
	ID       int8
	GioiTinh string
}

var danhSachSinhVien = make(map[string]SinhVien)

func themSinhVien(sv SinhVien) {
	danhSachSinhVien[strconv.Itoa(int(sv.ID))] = sv
	fmt.Println("Đã thêm sinh viên:", sv)
}

func suaSinhVien(id string, sv SinhVien) {
	if _, exists := danhSachSinhVien[id]; exists {
		danhSachSinhVien[id] = sv
		fmt.Println("Đã sửa thông tin sinh viên:", sv)
	} else {
		fmt.Println("Sinh viên không tồn tại!")
	}
}

func xoaSinhVien(id string) {
	if _, exists := danhSachSinhVien[id]; exists {
		delete(danhSachSinhVien, id)
		fmt.Println("Đã xóa sinh viên có ID:", id)
	} else {
		fmt.Println("Sinh viên không tồn tại!")
	}
}

func timKiemSinhVien(id string) {
	if sv, exists := danhSachSinhVien[id]; exists {
		fmt.Println("Thông tin sinh viên:", sv)
	} else {
		fmt.Println("Sinh viên không tồn tại!")
	}
}

func xuatSinhVienTheoGioiTinh(gioiTinh string) {
	fmt.Printf("Danh sách sinh viên %s:\n", gioiTinh)
	for _, sv := range danhSachSinhVien {
		if sv.GioiTinh == gioiTinh {
			fmt.Println(sv)
		}
	}
}

func main() {
	// Thêm sinh viên
	themSinhVien(SinhVien{ID: 1, Ten: "Nguyen Van A", Tuoi: 20, GioiTinh: "Nam"})
	themSinhVien(SinhVien{ID: 2, Ten: "Tran Thi B", Tuoi: 21, GioiTinh: "Nu"})

	// Sửa thông tin sinh viên
	suaSinhVien("1", SinhVien{ID: 1, Ten: "Nguyen Van A", Tuoi: 22, GioiTinh: "Nam"})

	// Tìm kiếm sinh viên
	timKiemSinhVien("2")

	// Xóa sinh viên
	xoaSinhVien("2")

	// Xuất sinh viên theo giới tính
	xuatSinhVienTheoGioiTinh("Nam")
	xuatSinhVienTheoGioiTinh("Nu")
}
