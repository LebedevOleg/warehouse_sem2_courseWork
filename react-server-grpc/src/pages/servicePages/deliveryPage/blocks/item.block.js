import {
	Box,
	FormControl,
	InputLabel,
	MenuItem,
	Select,
	TextField,
} from "@mui/material";
import React, { useContext } from "react";
import { itemContext } from "../item.Context";

const DeliveryItemBlock = (props) => {
	const { selectedItems, setSelectedItems } = useContext(itemContext);
	const [item, setItem] = React.useState({
		name: "",
		count: 0,
		price: 0,
	});

	const handleChangeItem = (e) => {
		setItem({
			...item,
			[e.target.name]: e.target.value,
		});
	};

	return (
		<>
			<Box>
				<FormControl>
					<InputLabel>Товар</InputLabel>
					<Select
						label="Товар"
						id="item"
						name="name"
						onChange={handleChangeItem}
					>
						{props.items.map((item) => (
							<MenuItem value={item.id}>{item.name}</MenuItem>
						))}
					</Select>
				</FormControl>
				<TextField
					type="number"
					label="Количество"
					name="count"
				></TextField>
			</Box>
		</>
	);
};

export default DeliveryItemBlock;
