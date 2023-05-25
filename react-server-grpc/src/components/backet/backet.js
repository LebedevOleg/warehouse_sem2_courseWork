import {
	Button,
	Dialog,
	DialogActions,
	DialogContent,
	DialogTitle,
	IconButton,
} from "@mui/material";
import React from "react";
import ShoppingCartIcon from "@mui/icons-material/ShoppingCart";

const BacketModal = () => {
	const [open, setOpen] = React.useState(false);
	let b;
	if (localStorage.getItem("backet") === null) {
		b = [];
		localStorage.setItem("backet", JSON.stringify([]));
	} else {
		b = JSON.parse(localStorage.getItem("backet"));
	}
	const backet = b;
	const handleOpen = () => setOpen(true);
	const handleClose = () => setOpen(false);

	const handleClick = () => {
		window.location.replace("/backet");
	};
	if (backet.length === 0) {
		<div>
			<IconButton aria-label="backet" onClick={handleOpen}>
				<ShoppingCartIcon />
			</IconButton>
			<Dialog open={open} onClose={handleClose}>
				<DialogTitle>Корзина</DialogTitle>
				<DialogContent>Нет товаров в корзине</DialogContent>
				<DialogActions>
					<Button onClick={handleClose}>Закрыть</Button>
					<Button onClick={handleClick}>Оформить</Button>
				</DialogActions>
			</Dialog>
		</div>;
	}

	return (
		<div>
			<IconButton aria-label="backet" onClick={handleOpen}>
				<ShoppingCartIcon />
			</IconButton>
			<Dialog open={open} onClose={handleClose}>
				<DialogTitle>Корзина</DialogTitle>
				<DialogContent>
					{backet.map((item) => (
						<div key={item.id}>
							{item.name}: {item.count} шт.
						</div>
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
