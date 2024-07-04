"use client"

import { useFrame } from '@react-three/fiber'
import React, { useRef } from 'react'
import { useGLTF } from '@react-three/drei'

export function Model(props) {
  const { nodes, materials } = useGLTF('model.glb')
  const groupRef = useRef(null)
  useFrame(() => {
    groupRef.current.rotation.y += 0.007
  })
  return (
    <group ref={groupRef} {...props} dispose={null} scale={[1.8, 1.8, 1.8]}>
      <mesh
        castShadow
        receiveShadow
        geometry={nodes.Cube.geometry}
        material={materials['Material.001']}
        position={[0, 0, 0]}
      />
    </group>
  )
}

useGLTF.preload('model.glb')

