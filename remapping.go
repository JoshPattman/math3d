package math3d

// AxesRemapper allows remapping of vectors from one axis space to another
type AxesRemapper struct {
	SrcForward Vec `json:"src_forward"`
	SecUp      Vec `json:"src_up"`
	SrcLeft    Vec `json:"src_left"`

	TarForward Vec `json:"target_forward"`
	TarUp      Vec `json:"target_up"`
	TarLeft    Vec `json:"target_left"`
}

// NewAxesRemapper creates an AxesRemapper that converts from src to tar
func NewAxesRemapper(SrcForward, SrcUp, SrcLeft, TarForward, TarUp, TarLeft Vec) *AxesRemapper {
	return &AxesRemapper{
		SrcForward, SrcUp, SrcLeft,
		TarForward, TarUp, TarLeft,
	}
}

// Remap the vector from coordinate system with Src axes to coordinate system with Target axes, whilst keeping the meaning of the vector
func (a *AxesRemapper) RemapVec(v Vec) Vec {
	fwd := v.Dot(a.SrcForward)
	lft := v.Dot(a.SrcLeft)
	up := v.Dot(a.SecUp)
	return a.TarForward.Mul(fwd).Add(a.TarLeft.Mul(lft)).Add(a.TarUp.Mul(up))
}

// Remap the vector from coordinate system with Target axes to coordinate system with Src axes, whilst keeping the meaning of the vector
func (a *AxesRemapper) RemapVecInverse(v Vec) Vec {
	fwd := v.Dot(a.TarForward)
	lft := v.Dot(a.TarLeft)
	up := v.Dot(a.TarUp)
	return a.SrcForward.Mul(fwd).Add(a.SrcLeft.Mul(lft)).Add(a.SecUp.Mul(up))
}

/*func (a *AxesRemapper) RemapQuat(q Quat) Quat {
	return q
}*/
