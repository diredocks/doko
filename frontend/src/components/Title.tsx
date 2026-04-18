import { Title as SolidTitle } from "@solidjs/meta";

export default function Title(props: { children: string }) {
	return <SolidTitle>{`Doko - ${props.children}`}</SolidTitle>;
}
