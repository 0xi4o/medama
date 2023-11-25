import { client, type ClientOptions, type DataResponse } from './client';

const statsSummary = async (
	opts: ClientOptions<'StatsSummary'>
): Promise<DataResponse<'StatsSummary'>> => {
	const res = await client('/website/{hostname}/summary', {
		method: 'GET',
		...opts,
	});
	return { data: await res.json(), res };
};

const statsPages = async (
	opts: ClientOptions<'StatsPages'>
): Promise<DataResponse<'StatsPages'>> => {
	const res = await client('/website/{hostname}/pages', {
		method: 'GET',
		...opts,
	});
	return { data: await res.json(), res };
};

const statsTime = async (
	opts: ClientOptions<'StatsTime'>
): Promise<DataResponse<'StatsTime'>> => {
	const res = await client('/website/{hostname}/time', {
		method: 'GET',
		...opts,
	});
	return { data: await res.json(), res };
};

export { statsPages, statsSummary, statsTime };
