import { Button, Typography } from "@mui/material";
import axios from "axios";
import React, { useCallback, useEffect } from "react";
import BacketItem from "./backetItem";

const BacketPage = () => {
	const [backet, setBacket] = React.useState([]);

	const getBacket = async () => {
		setBacket(localStorage.getItem("backet"));
	};

	const handleOffer = () => {
		setBacket(localStorage.getItem("backet"));
		axios.post("http://localhost:8000/offer", {});
	};

	useEffect(() => {
		getBacket();
	}, []);
	return (
		<div>
			<Typography variant="h3">Корзина</Typography>
			<div>
				{backet.map((item) => (
					<div key={item.id}>
						<BacketItem item={item} />
					</div>
				))}
			</div>
			<Button onClick={handleOffer}>Оформить</Button>
		</div>
	);
};

export default BacketPage;
