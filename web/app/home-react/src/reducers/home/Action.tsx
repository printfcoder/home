export const CHANGE_BOTTOM_TAB = 'CHANGE_BOTTOM_TAB';

export function updateCart(tabName: string) {
    return {
        type: CHANGE_BOTTOM_TAB,
        payload: {
            tabName,
        }
    }
}