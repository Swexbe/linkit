import LinkIcon from "@material-ui/icons/Link";
import DeleteIcon from "@material-ui/icons/Delete";

const LinkItem = (link, onDelete) => {
    return {
        id: link.shortcut,
        icon: LinkIcon,
        text: link.shortcut,
        secondaryText: link.linkurl,
        actionIcon: DeleteIcon,
        actionOnClick: onDelete,
    };
};

export default LinkItem;