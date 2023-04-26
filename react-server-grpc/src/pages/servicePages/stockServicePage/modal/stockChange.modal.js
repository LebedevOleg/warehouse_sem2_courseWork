import {
	Box,
	Button,
	Dialog,
	DialogActions,
	DialogContent,
	DialogTitle,
	TextField,
} from "@mui/material";
import axios from "axios";
import React, { useState } from "react";
import { useAuth } from "../../../../hooks/auth.hook";

const StockChangeModal = (props) => {
	const [openModal, setOpenModal] = useState(false);
	const { token } = useAuth();
	const [stock, setStock] = useState({
		id: props.stock.id,
		name: props.stock.name,
		address: props.stock.address,
	});

	const handleOpenModal = () => {
		setOpenModal(true);
	};
	const handleCloseModal = () => {
		setOpenModal(false);
	};
	const handleChangeStock = (e) => {
		setStock({
			...stock,
			[e.target.name]: e.target.value,
		});
	};
	const handleUpdateStock = async (e) => {
		// todo: check work
		await axios
			.post("http://localhost:8000/updateStock", stock)
			.then((res) => {
				setOpenModal(false);
			});
	};

	return (
		<div>
			<Button onClick={handleOpenModal}>Изменить</Button>
			<Dialog open={openModal} onClose={handleCloseModal} fullWidth>
				<DialogTitle>Изменить склад</DialogTitle>
				<DialogContent>
					<Box
						sx={{
							display: "flex",
							alignItems: "flex-end",
						}}
					>
						<TextField
							label="Название склада"
							name="name"
							defaultValue={props.stock.name}
							onChange={handleChangeStock}
							margin="dense"
							variant="filled"
							fullWidth
						/>
					</Box>
					<Box sx={{ display: "flex", alignItems: "flex-end" }}>
						<TextField
							label="Адрес склада"
							name="address"
							defaultValue={props.stock.address}
							onChange={handleChangeStock}
							margin="dense"
							variant="filled"
							fullWidth
						/>
					</Box>
				</DialogContent>
				<DialogActions>
					<Button onClick={handleCloseModal}>Отмена</Button>
					<Button onClick={handleUpdateStock}>Изменить</Button>
				</DialogActions>
			</Dialog>
		</div>
	);
};

export default StockChangeModal;
