import {
	Button,
	Dialog,
	DialogActions,
	DialogContent,
	DialogTitle,
	IconButton,
} from "@mui/material";
import React, { useContext } from "react";
import ShoppingCartIcon from "@mui/icons-material/ShoppingCart";

const BacketModal = () => {
	const [open, setOpen] = React.useState(false);
	const backet = JSON.parse(localStorage.getItem("backet"));

	const handleOpen = () => setOpen(true);
	const handleClose = () => setOpen(false);

	const handleClick = () => {
		window.location.replace("/backet");
	};

	return (
		<div>
			<IconButton aria-label="backet" onClick={handleOpen}>
				<ShoppingCartIcon />
			</IconButton>
			<Dialog open={open} onClose={handleClose}>
				<DialogTitle>Корзина</DialogTitle>
				<DialogContent>
					{backet.map((item) => (
						<div key={item.id}>{item.name}</div>
					))}
				</DialogContent>
				<DialogActions>
					<Button onClick={handleClose}>Закрыть</Button>
					<Button onClick={handleClick}>Оформить</Button>
				</DialogActions>
			</Dialog>
		</div>
	);
};

export default BacketModal;
