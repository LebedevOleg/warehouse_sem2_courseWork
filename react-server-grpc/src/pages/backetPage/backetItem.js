import { Button, Typography } from "@mui/material";
import React from "react";

const BacketItem = ({ item }) => {
	const deleteItem = (id) => {
		let backet = JSON.parse(localStorage.getItem("backet"));
		backet = backet.filter((item) => item.id !== id);
		localStorage.setItem("backet", JSON.stringify(backet));
	};

	const addItem = (id) => {
		let backet = JSON.parse(localStorage.getItem("backet"));
		backet.map((item) => {
			if (item.id === id) {
				item.count++;
			}
		});
		localStorage.setItem("backet", JSON.stringify(backet));
	};

	const minusItem = (id) => {
		let backet = JSON.parse(localStorage.getItem("backet"));
		backet.map((item) => {
			if (item.id === id) {
				item.count--;
				if (item.count <= 0) {
					deleteItem(item.id);
				}
			}
		});
		localStorage.setItem("backet", JSON.stringify(backet));
	};

	return (
		<div key={item.id}>
			<Typography>{item.name}</Typography>
			<Button onClick={() => minusItem(item.id)}>-</Button>
			<Typography>{item.count}</Typography>
			<Button onClick={() => addItem(item.id)}>+</Button>
			<Typography>{item.dim}</Typography>
			<Typography>{item.price * item.count}</Typography>
			<Button onClick={() => deleteItem(item.id)}>Удалить</Button>
		</div>
	);
};

export default BacketItem;
