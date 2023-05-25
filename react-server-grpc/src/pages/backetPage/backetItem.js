import { Button, Typography } from "@mui/material";
import React from "react";

const BacketItem = ({ item, update }) => {
	const [data, setData] = React.useState(item);
	const deleteItem = (id) => {
		let backet = JSON.parse(localStorage.getItem("backet"));
		backet = backet.filter((i) => i.id !== id);
		localStorage.setItem("backet", JSON.stringify(backet));
		update();
	};

	const addItem = (id) => {
		let backet = JSON.parse(localStorage.getItem("backet"));
		backet.map((i) => {
			if (i.id === id) {
				i.count++;
				setData(i);
			}
		});
		localStorage.setItem("backet", JSON.stringify(backet));
	};

	const minusItem = (id) => {
		let backet = JSON.parse(localStorage.getItem("backet"));
		backet.map((i) => {
			if (i.id === id) {
				i.count--;
				if (i.count <= 0) {
					deleteItem(i.id);
				} else {
					setData(i);
				}
			}
		});
		localStorage.setItem("backet", JSON.stringify(backet));
	};

	return (
		<div key={item.id}>
			<Typography>{data.name}</Typography>
			<Button onClick={() => minusItem(item.id)}>-</Button>
			<Typography>{data.count}</Typography>
			<Button onClick={() => addItem(item.id)}>+</Button>
			<Typography>{data.dim}</Typography>
			<Typography>{data.price * data.count}</Typography>
			<Button onClick={() => deleteItem(item.id)}>Удалить</Button>
		</div>
	);
};

export default BacketItem;
