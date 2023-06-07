import {
	Box,
	Button,
	FormControl,
	InputLabel,
	MenuItem,
	Select,
	Typography,
} from "@mui/material";
import React, { useCallback, useContext, useEffect } from "react";
import { itemContext } from "./item.Context";
import axios from "axios";
import { AuthContext } from "../../../context/auth.context";
import FullFeaturedCrudGrid from "./blocks/item.block";
import Docxtemplater from "docxtemplater";
import JSZip from "jszip";
import { saveAs } from "file-saver";

const DeliveryPage = () => {
	const auth = useContext(AuthContext);
	const [deliveries, setDeliveries] = React.useState({
		provider: 0,
		storage: 0,
	});
	const [providers, setProviders] = React.useState([]);
	const [storages, setStorages] = React.useState([]);

	const [items, setItems] = React.useState([]);
	const [selectedItems, setSelectedItems] = React.useState([]);

	const handleChangeDeliver = (e) => {
		setDeliveries({
			...deliveries,
			[e.target.name]: e.target.value,
		});
	};
	const GetAllProviders = useCallback(async () => {
		await axios
			.get("http://localhost:8000/allproviders", {
				headers: { Authorization: `Bearer ${auth.token}` },
			})
			.then((res) => {
				setProviders(res.data.providers);
			});
	}, []);
	const GetAllItems = useCallback(async () => {
		await axios.get("http://localhost:8000/getallitems").then((res) => {
			setItems(res.data.allItems);
		});
	}, []);
	const GetAllStorages = useCallback(async () => {
		await axios
			.get("http://localhost:8000/allstorages", {
				headers: { Authorization: `Bearer ${auth.token}` },
			})
			.then((res) => {
				setStorages(res.data.storages);
			});
	}, []);

	const handleCreateDelivery = async () => {
		await axios
			.post(
				"http://localhost:8000/addtransaction",
				{
					provider: deliveries.provider,
					storage: deliveries.storage,
					items: selectedItems,
				},
				{
					headers: { Authorization: `Bearer ${auth.token}` },
				}
			)
			.then((res) => {
				fetch("http://localhost:8000/getfile", {
					headers: { Authorization: `Bearer ${auth.token}` },
				})
					.then((responce) => responce.arrayBuffer())
					.then((buffer) => {
						console.log(res.data.data);

						let doc = new Docxtemplater();
						doc.loadZip(new JSZip(buffer));
						doc.setData(res.data.data);
						doc.render();
						let out = doc.getZip().generate({
							type: "blob",
							mimeType:
								"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
						});
						saveAs(out, "Накладная.docx");
					});
			});
	};

	useEffect(() => {
		GetAllProviders();
		GetAllItems();
		GetAllStorages();
	}, [GetAllProviders, GetAllItems, GetAllStorages]);

	return (
		<div>
			<Typography variant="h3">Сформировать накладную</Typography>
			<Box>
				<FormControl sx={{ mr: 1 }}>
					<InputLabel>Поставщик</InputLabel>
					<Select
						sx={{ width: "250px" }}
						label="Поставщик"
						id="provider"
						name="provider"
						onChange={handleChangeDeliver}
					>
						{providers.map((provider) => (
							<MenuItem value={provider.id} key={provider.id}>
								{provider.name}
							</MenuItem>
						))}
					</Select>
				</FormControl>
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
			</Box>
			<Button>Добавить нового поставщика</Button>
			<itemContext.Provider value={[selectedItems, setSelectedItems]}>
				<FullFeaturedCrudGrid
					items={items.sort((a, b) => a.id - b.id)}
				/>
			</itemContext.Provider>
			<Box>
				<Button onClick={handleCreateDelivery}>
					Сформировать накладную
				</Button>
			</Box>
		</div>
	);
};

export default DeliveryPage;
