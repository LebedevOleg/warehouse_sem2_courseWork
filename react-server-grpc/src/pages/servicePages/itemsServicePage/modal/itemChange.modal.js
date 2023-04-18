import {
	Button,
	Dialog,
	DialogActions,
	DialogContent,
	DialogTitle,
	FormControl,
	InputLabel,
	MenuItem,
	Select,
	TextField,
} from "@mui/material";
import { Box } from "@mui/system";
import axios from "axios";
import React, { useCallback, useEffect, useState } from "react";
import { useAuth } from "../../../../hooks/auth.hook";

const ItemChangeModal = (props) => {
	const [openModal, setOpenModal] = useState(false);
	const { token } = useAuth();
	const [item, setItem] = useState({
		id: props.item.id,
		name: props.item.name,
		pfu: props.item.pfu,
		desc: props.item.desc,
		dim: props.item.dim,
		category: props.item.category,
	});
	const [category, setCategory] = useState([]);
	const [categoryId, setCategoryId] = useState(0);
	const GetCategories = useCallback(async () => {
		await axios
			.get("http://localhost:8000/getallcategories")
			.then((res) => {
				setCategory(res.data.categories);
			});
	}, []);
	useEffect(() => {
		GetCategories();
	}, [GetCategories]);

	const handleChangeItem = (e) => {
		if (e.target.name === "pfu") {
			setItem({
				...item,
				[e.target.name]: parseFloat(e.target.value),
			});
			return;
		}
		setItem({
			...item,
			[e.target.name]: e.target.value,
		});
	};

	const handleChangeCategory = (e) => {
		setItem({
			...item,
			[e.target.name]: e.target.value,
		});
		setCategoryId(e.target.value);
	};

	const handleOpenModal = () => {
		GetCategories();
		setOpenModal(true);
	};
	const handleCloseModal = () => {
		setOpenModal(false);
	};

	const handleUpdateItem = async (e) => {
		await axios
			.post(
				"http://localhost:8000/updateitem",
				{ ...item },
				{ headers: { Authorization: `Bearer ${token}` } }
			)
			.then((res) => {});
	};

	return (
		<>
			<Button onClick={handleOpenModal}>Изменить</Button>
			<Dialog open={openModal} onClose={handleCloseModal} fullWidth>
				<DialogTitle>Изменить Товар</DialogTitle>
				<DialogContent>
					<Box
						sx={{
							display: "flex",
							alignItems: "flex-end",
						}}
					>
						<TextField
							defaultValue={props.item.name}
							name="name"
							margin="dense"
							id="name"
							label="Название предмета"
							fullWidth
							variant="filled"
							onChange={handleChangeItem}
						/>
					</Box>
					<Box
						sx={{
							display: "flex",
							alignItems: "flex-end",
						}}
					>
						<TextField
							defaultValue={props.item.dim}
							name="dim"
							margin="dense"
							id="dim"
							label="Измерение колличества предмета"
							fullWidth
							variant="filled"
							onChange={handleChangeItem}
						/>
					</Box>
					<Box
						sx={{
							display: "flex",
							alignItems: "flex-end",
						}}
					>
						<TextField
							defaultValue={props.item.pfu}
							name="pfu"
							margin="dense"
							id="pfu"
							type="number"
							label="Цена за единицу ппедмета"
							fullWidth
							variant="filled"
							onChange={handleChangeItem}
						/>
					</Box>
					<FormControl fullWidth>
						<InputLabel id="demo-simple-select-label">
							Категория
						</InputLabel>
						<Select
							labelId="category"
							id="category"
							defaultValue={props.item.c_id}
							label="Категория"
							name="category"
							onChange={handleChangeCategory}
						>
							{category.map((item) => (
								<MenuItem value={item.id}>{item.name}</MenuItem>
							))}
						</Select>
					</FormControl>
					<Box
						sx={{
							display: "flex",
							alignItems: "flex-end",
						}}
					>
						<TextField
							defaultValue={props.item.desc}
							name="desc"
							margin="dense"
							id="desc"
							label="Описание предмета"
							fullWidth
							variant="filled"
							onChange={handleChangeItem}
						/>
					</Box>
				</DialogContent>
				<DialogActions>
					<Button onClick={handleUpdateItem}>
						Сохранить изменения
					</Button>
					<Button onClick={handleCloseModal}>Отмена</Button>
				</DialogActions>
			</Dialog>
		</>
	);
};

export default ItemChangeModal;
