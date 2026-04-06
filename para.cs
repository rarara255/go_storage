using UnityEngine;

public class PlayerController : MonoBehavior
{
    public float speed = 5f;
    private Rigidbody2D rd;
    private Vector2 movement;

    void Start()
    {
        rd = GetComponent<Rigidbody2D>();
    }

    void Update()
    {
        movement.x = input.GetAxis("Horizontal");
    }

    void FixedUpdate()
    {
        rd.linearVelocity = new Vector2(movement.x * speed, rd.linearVelocity.y)
    }
}