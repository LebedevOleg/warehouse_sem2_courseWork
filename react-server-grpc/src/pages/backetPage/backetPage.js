import {
	Button,
	FormControl,
	InputLabel,
	MenuItem,
	Select,
	Typography,
} from "@mui/material";
import axios from "axios";
import React, { useCallback, useContext, useEffect } from "react";
import { AuthContext } from "../../context/auth.context";
import BacketItem from "./backetItem";

const BacketPage = () => {
	const auth = useContext(AuthContext);
	let b;
	if (localStorage.getItem("backet") === null) {
		b = [];
		localStorage.setItem("backet", JSON.stringify([]));
	} else {
		b = JSON.parse(localStorage.getItem("backet"));
	}
	const [backet, setBacket] = React.useState(b);
	const [storages, setStorages] = React.useState([]);
	const [storage, setStorage] = React.useState(0);

	const updateBacket = () => {
		if (localStorage.getItem("backet") === null) {
			b = [];
			localStorage.setItem("backet", JSON.stringify([]));
		} else {
			b = JSON.parse(localStorage.getItem("backet"));
		}
		setBacket(b);
	};
	const handleChangeDeliver = (e) => {
		setStorage(e.target.value);
	};

	const getStorages = useCallback(async () => {
		await axios
			.get("http://localhost:8000/allstorages", {
				headers: { Authorization: `Bearer ${auth.token}` },
			})
			.then((res) => {
				setStorages(res.data.storages);
			});
	}, []);

	const handleOffer = () => {
		updateBacket();
		axios.post(
			"http://localhost:8000/createoffer",
			{
				storage_id: storage,
				items: backet,
			},
			{
				headers: { Authorization: `Bearer ${auth.token}` },
			}
		);
	};

	useEffect(() => {
		getStorages();
	}, [getStorages]);

	return (
		<div>
			<Typography variant="h3">Корзина</Typography>
			<FormControl>
				<InputLabel>Склад</InputLabel>
				<Select
					sx={{ width: "250px" }}
					label="Склад"
					id="storage"
					name="storage"
					onChange={handleChangeDeliver}
				>
					{storages.map((storage) => (
						<MenuItem value={storage.id} key={storage.id}>
							{storage.name}
						</MenuItem>
					))}
				</Select>
			</FormControl>
			<div>
				{backet.map((item) => (
					<div key={item.id}>
						<BacketItem item={item} update={updateBacket} />
					</div>
				))}
			</div>
			{(backet.length > 0 && (
				<Button onClick={handleOffer}>Оформить</Button>
			)) || <Button disabled>Оформить</Button>}
		</div>
	);
};

export default BacketPage;
