import { Stack, Typography } from "@mui/material";
import axios from "axios";
import React, { useCallback, useContext } from "react";
import { AuthContext } from "../../../context/auth.context";
import OrderBlock from "./orderBlock";

const OrdersPage = () => {
	const [orders, setOrders] = React.useState([]);
	const auth = useContext(AuthContext);

	const getOrders = useCallback(async () => {
		await axios
			.get("http://localhost:8000/getallorders", {
				headers: { Authorization: `Bearer ${auth.token}` },
			})
			.then((res) => {
				setOrders(res.data.orders);
			});
	}, []);

	React.useEffect(() => {
		getOrders();
	}, [getOrders]);

	return (
		<>
			<Typography variant="h3">Заказы</Typography>
			<Stack spacing={2} direction="row">
				<Typography>ID заказа</Typography>
				<Typography>Дата Начала</Typography>
				<Typography>Дата завершения</Typography>
				<Typography>Статус</Typography>
				<Typography>Цена</Typography>
				<Typography>Адрес Склада</Typography>
				<Typography>ID пользователя</Typography>
			</Stack>

			{orders.map((order) => (
				<OrderBlock order={order} />
			))}
		</>
	);
};

export default OrdersPage;
