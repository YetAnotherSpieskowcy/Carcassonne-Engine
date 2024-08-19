__all__ = ("format_binary_tile_bits",)


_FEATURE_BIT_SIZE = 10
_MODIFIER_BIT_SIZE = 4
_MEEPLE_BIT_SIZE = 9
# from LSB to MSB (excluding last group)
_GROUP_SIZES = (
    _FEATURE_BIT_SIZE,
    _FEATURE_BIT_SIZE,
    _FEATURE_BIT_SIZE,
    _MODIFIER_BIT_SIZE,
    2,
    _MEEPLE_BIT_SIZE,
)


def format_binary_tile_bits(bits: int) -> str:
    """
    Utility for "pretty" representation of a binary tile integer
    grouped by how they're interpreted.
    """
    raw_bits = f"{bits:064b}"

    groups = []
    end = len(raw_bits)
    for group_size in _GROUP_SIZES:
        start = end - group_size
        groups.append(raw_bits[start:end])
        end = start
    groups.append(raw_bits[:end])
    groups.reverse()

    return "_".join(groups)
